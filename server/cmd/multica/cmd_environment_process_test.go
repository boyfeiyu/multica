package main

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"
)

func TestEnvironmentCLIExitCode(t *testing.T) {
	binary := filepath.Join(t.TempDir(), "multica")
	if runtime.GOOS == "windows" {
		binary += ".exe"
	}
	build := exec.Command(filepath.Join(runtime.GOROOT(), "bin", "go"), "build", "-o", binary, ".")
	build.Dir = testLaunchDir
	if output, err := build.CombinedOutput(); err != nil {
		t.Fatalf("build CLI: %v\n%s", err, output)
	}
	data, err := os.ReadFile(binary)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("production binary SHA256=%x", sha256.Sum256(data))
	const rejection = "environment commands need the credential of a task running in this daemon; run them from inside the task, while it is running\n"
	for _, tc := range []struct {
		name      string
		status    int
		body      string
		wantExit  int
		locale    string
		debugFlag bool
		debugEnv  string
	}{
		{name: "empty", status: http.StatusOK, body: `{"environments":[]}`},
		{name: "running", status: http.StatusOK, body: `{"environments":[{"name":"web","state":"running"}]}`},
		{name: "rejected task", status: http.StatusUnauthorized, body: rejection, wantExit: 3},
		{name: "debug flag", status: http.StatusUnauthorized, body: rejection, wantExit: 3, debugFlag: true},
		{name: "debug environment", status: http.StatusUnauthorized, body: rejection, wantExit: 3, debugEnv: "1"},
		{name: "Chinese", status: http.StatusUnauthorized, body: rejection, wantExit: 3, locale: "zh_CN.UTF-8"},
		{name: "Chinese debug", status: http.StatusUnauthorized, body: rejection, wantExit: 3, locale: "zh_CN.UTF-8", debugFlag: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			requests := make(chan string, 1)
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				body, _ := io.ReadAll(r.Body)
				requests <- fmt.Sprintf("%s %s Authorization=%q body=%s", r.Method, r.URL.Path, r.Header.Get("Authorization"), body)
				w.WriteHeader(tc.status)
				fmt.Fprint(w, tc.body)
			}))
			defer server.Close()
			u, _ := url.Parse(server.URL)
			t.Setenv("MULTICA_DAEMON_PORT", u.Port())
			t.Setenv("MULTICA_TOKEN", "mat_SYNTHETIC_ENVIRONMENT_TEST")
			locale := tc.locale
			if locale == "" {
				locale = "en_US.UTF-8"
			}
			t.Setenv("LC_ALL", locale)
			t.Setenv("MULTICA_DEBUG", tc.debugEnv)
			ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
			defer cancel()
			args := []string{"env", "list", "--output", "json"}
			if tc.debugFlag {
				args = append(args, "--debug")
			}
			child := exec.CommandContext(ctx, binary, args...)
			child.Dir, err = os.Getwd()
			if err != nil {
				t.Fatal(err)
			}
			var stdout, stderr bytes.Buffer
			child.Stdout, child.Stderr = &stdout, &stderr
			err = child.Run()
			if ctx.Err() != nil {
				t.Fatal(ctx.Err())
			}
			if child.ProcessState == nil {
				t.Fatalf("start CLI: %v", err)
			}
			exit := child.ProcessState.ExitCode()
			t.Logf("status=%d exit=%d stdout=%q stderr=%q", tc.status, exit, stdout.String(), stderr.String())
			select {
			case request := <-requests:
				t.Log(request)
				if !strings.Contains(request, `POST /environments/list Authorization="Bearer mat_SYNTHETIC_ENVIRONMENT_TEST" body={}`) {
					t.Errorf("request = %s", request)
				}
			default:
				t.Fatal("CLI did not reach the synthetic daemon")
			}
			if exit != tc.wantExit {
				t.Errorf("exit = %d, want %d", exit, tc.wantExit)
			}
			if tc.wantExit != 0 && stdout.Len() != 0 {
				t.Errorf("error stdout = %q, want empty", stdout.String())
			}
			if tc.wantExit == 0 {
				var environments []daemonEnvironment
				if err := json.Unmarshal(stdout.Bytes(), &environments); err != nil {
					t.Fatalf("decode successful CLI output: %v", err)
				}
				if tc.name == "empty" && len(environments) != 0 {
					t.Errorf("empty list = %v", environments)
				}
				if tc.name == "running" && (len(environments) != 1 || environments[0].Name != "web" || environments[0].State != "running") {
					t.Errorf("running list = %v", environments)
				}
				if stderr.Len() != 0 {
					t.Errorf("success stderr = %q", stderr.String())
				}
				return
			}
			if !strings.Contains(stderr.String(), strings.TrimSpace(rejection)) {
				t.Errorf("stderr lost daemon guidance: %q", stderr.String())
			}
			stop := "Stop here. Do not retry or fall back to a profile or member credential."
			if tc.locale != "" {
				stop = "请到此为止。不要重试，也不要改用 profile 或成员凭证。"
			}
			if !strings.Contains(stderr.String(), stop) || strings.Contains(stderr.String(), "multica login") {
				t.Errorf("unsafe task rejection guidance: %q", stderr.String())
			}
			debug := tc.debugFlag || tc.debugEnv != ""
			if strings.Contains(stderr.String(), "[debug] http:") != debug {
				t.Errorf("debug HTTP detail = %q, debug=%v", stderr.String(), debug)
			}
		})
	}
}

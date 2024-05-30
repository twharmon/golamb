package golamb

import (
	"log"
	"os"
	"testing"
)

func TestContextLogDebug(t *testing.T) {
	fname := "TestContextLogDebug"
	f, err := os.Create(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		f.Close()
		os.Remove(fname)
	}()
	want := "[DEBUG] foo bar\n"
	os.Stdout = f
	ctx := &handlerContext{
		logger:   NewDefaultLogger(),
		logLevel: LogLevelDebug,
	}
	ctx.LogDebug("foo %s", "bar")
	b, err := os.ReadFile(fname)
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
	}
	got := string(b)
	if want != got {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestContextLogInfo(t *testing.T) {
	fname := "TestContextLogInfo"
	f, err := os.Create(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		f.Close()
		os.Remove(fname)
	}()
	want := "[INFO] foo bar\n"
	os.Stdout = f
	ctx := &handlerContext{
		logger:   NewDefaultLogger(),
		logLevel: LogLevelInfo,
	}
	ctx.LogInfo("foo %s", "bar")
	b, err := os.ReadFile(fname)
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
	}
	got := string(b)
	if want != got {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestContextLogNotice(t *testing.T) {
	fname := "TestContextLogNotice"
	f, err := os.Create(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		f.Close()
		os.Remove(fname)
	}()
	want := "[NOTICE] foo bar\n"
	os.Stdout = f
	ctx := &handlerContext{
		logger:   NewDefaultLogger(),
		logLevel: LogLevelNotice,
	}
	ctx.LogNotice("foo %s", "bar")
	b, err := os.ReadFile(fname)
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
	}
	got := string(b)
	if want != got {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestContextLogWarning(t *testing.T) {
	fname := "TestContextLogWarning"
	f, err := os.Create(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		f.Close()
		os.Remove(fname)
	}()
	want := "[WARNING] foo bar\n"
	os.Stdout = f
	ctx := &handlerContext{
		logger:   NewDefaultLogger(),
		logLevel: LogLevelWarning,
	}
	ctx.LogWarning("foo %s", "bar")
	b, err := os.ReadFile(fname)
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
	}
	got := string(b)
	if want != got {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestContextLogError(t *testing.T) {
	fname := "TestContextLogError"
	f, err := os.Create(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		f.Close()
		os.Remove(fname)
	}()
	want := "[ERROR] foo bar\n"
	os.Stdout = f
	ctx := &handlerContext{
		logger:   NewDefaultLogger(),
		logLevel: LogLevelError,
	}
	ctx.LogError("foo %s", "bar")
	b, err := os.ReadFile(fname)
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
	}
	got := string(b)
	if want != got {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestContextLogAlert(t *testing.T) {
	fname := "TestContextLogAlert"
	f, err := os.Create(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		f.Close()
		os.Remove(fname)
	}()
	want := "[ALERT] foo bar\n"
	os.Stdout = f
	ctx := &handlerContext{
		logger:   NewDefaultLogger(),
		logLevel: LogLevelAlert,
	}
	ctx.LogAlert("foo %s", "bar")
	b, err := os.ReadFile(fname)
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
	}
	got := string(b)
	if want != got {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestContextLogCritical(t *testing.T) {
	fname := "TestContextLogCritical"
	f, err := os.Create(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		f.Close()
		os.Remove(fname)
	}()
	want := "[CRITICAL] foo bar\n"
	os.Stdout = f
	ctx := &handlerContext{
		logger:   NewDefaultLogger(),
		logLevel: LogLevelCritical,
	}
	ctx.LogCritical("foo %s", "bar")
	b, err := os.ReadFile(fname)
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
	}
	got := string(b)
	if want != got {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestContextLogEmergency(t *testing.T) {
	fname := "TestContextLogEmergency"
	f, err := os.Create(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		f.Close()
		os.Remove(fname)
	}()
	want := "[EMERGENCY] foo bar\n"
	os.Stdout = f
	ctx := &handlerContext{
		logger:   NewDefaultLogger(),
		logLevel: LogLevelEmergency,
	}
	ctx.LogEmergency("foo %s", "bar")
	b, err := os.ReadFile(fname)
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
	}
	got := string(b)
	if want != got {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestContextLogSilent(t *testing.T) {
	fname := "TestContextLogSilent"
	f, err := os.Create(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		f.Close()
		os.Remove(fname)
	}()
	want := ""
	os.Stdout = f
	ctx := &handlerContext{
		logger:   NewDefaultLogger(),
		logLevel: LogLevelSilent,
	}
	ctx.LogEmergency("foo %s", "bar")
	b, err := os.ReadFile(fname)
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
	}
	got := string(b)
	if want != got {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestContextLogCriticalNone(t *testing.T) {
	fname := "TestContextLogCriticalNone"
	f, err := os.Create(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		f.Close()
		os.Remove(fname)
	}()
	want := ""
	os.Stdout = f
	ctx := &handlerContext{
		logger:   NewDefaultLogger(),
		logLevel: LogLevelSilent,
	}
	ctx.LogCritical("foo %s", "bar")
	b, err := os.ReadFile(fname)
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
	}
	got := string(b)
	if want != got {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestContextLogAlertNone(t *testing.T) {
	fname := "TestContextLogAlertNone"
	f, err := os.Create(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		f.Close()
		os.Remove(fname)
	}()
	want := ""
	os.Stdout = f
	ctx := &handlerContext{
		logger:   NewDefaultLogger(),
		logLevel: LogLevelSilent,
	}
	ctx.LogAlert("foo %s", "bar")
	b, err := os.ReadFile(fname)
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
	}
	got := string(b)
	if want != got {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestContextLogErrorNone(t *testing.T) {
	fname := "TestContextLogErrorNone"
	f, err := os.Create(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		f.Close()
		os.Remove(fname)
	}()
	want := ""
	os.Stdout = f
	ctx := &handlerContext{
		logger:   NewDefaultLogger(),
		logLevel: LogLevelSilent,
	}
	ctx.LogError("foo %s", "bar")
	b, err := os.ReadFile(fname)
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
	}
	got := string(b)
	if want != got {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestContextLogNoticeNone(t *testing.T) {
	fname := "TestContextLogNoticeNone"
	f, err := os.Create(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		f.Close()
		os.Remove(fname)
	}()
	want := ""
	os.Stdout = f
	ctx := &handlerContext{
		logger:   NewDefaultLogger(),
		logLevel: LogLevelSilent,
	}
	ctx.LogNotice("foo %s", "bar")
	b, err := os.ReadFile(fname)
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
	}
	got := string(b)
	if want != got {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestContextLogInfoNone(t *testing.T) {
	fname := "TestContextLogInfoNone"
	f, err := os.Create(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		f.Close()
		os.Remove(fname)
	}()
	want := ""
	os.Stdout = f
	ctx := &handlerContext{
		logger:   NewDefaultLogger(),
		logLevel: LogLevelSilent,
	}
	ctx.LogInfo("foo %s", "bar")
	b, err := os.ReadFile(fname)
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
	}
	got := string(b)
	if want != got {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestContextLogDebugNone(t *testing.T) {
	fname := "TestContextLogDebugNone"
	f, err := os.Create(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		f.Close()
		os.Remove(fname)
	}()
	want := ""
	os.Stdout = f
	ctx := &handlerContext{
		logger:   NewDefaultLogger(),
		logLevel: LogLevelSilent,
	}
	ctx.LogDebug("foo %s", "bar")
	b, err := os.ReadFile(fname)
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
	}
	got := string(b)
	if want != got {
		t.Fatalf("want %v; got %v", want, got)
	}
}

func TestContextLogWarningNone(t *testing.T) {
	fname := "TestContextLogWarningNone"
	f, err := os.Create(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		f.Close()
		os.Remove(fname)
	}()
	want := ""
	os.Stdout = f
	ctx := &handlerContext{
		logger:   NewDefaultLogger(),
		logLevel: LogLevelSilent,
	}
	ctx.LogWarning("foo %s", "bar")
	b, err := os.ReadFile(fname)
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
	}
	got := string(b)
	if want != got {
		t.Fatalf("want %v; got %v", want, got)
	}
}

package golamb

import (
	"io/ioutil"
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
	ctx := &handlerContext{}
	ctx.LogDebug("foo %s", "bar")
	b, err := ioutil.ReadFile(fname)
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
	ctx := &handlerContext{}
	ctx.LogInfo("foo %s", "bar")
	b, err := ioutil.ReadFile(fname)
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
	ctx := &handlerContext{}
	ctx.LogNotice("foo %s", "bar")
	b, err := ioutil.ReadFile(fname)
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
	ctx := &handlerContext{}
	ctx.LogWarning("foo %s", "bar")
	b, err := ioutil.ReadFile(fname)
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
	ctx := &handlerContext{}
	ctx.LogError("foo %s", "bar")
	b, err := ioutil.ReadFile(fname)
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
	ctx := &handlerContext{}
	ctx.LogAlert("foo %s", "bar")
	b, err := ioutil.ReadFile(fname)
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
	ctx := &handlerContext{}
	ctx.LogCritical("foo %s", "bar")
	b, err := ioutil.ReadFile(fname)
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
	ctx := &handlerContext{}
	ctx.LogEmergency("foo %s", "bar")
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
	}
	got := string(b)
	if want != got {
		t.Fatalf("want %v; got %v", want, got)
	}
}

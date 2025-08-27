package client_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"

	"terminal_chatroom/client"
)

func TestRun(t *testing.T) {
	input := strings.NewReader("hello world!\n/exit\n")
	output := &bytes.Buffer{}

	want := "Enter text:\nYou wrote: hello world!\nYou wrote: /exit\ngoodbye!\n"

	c := client.New(&client.Config{
		Host:   "localhost",
		Port:   "3333",
		Reader: input,
		Writer: output,
	})

	c.Run()

	got := output.String()

	if want != got {
		t.Error(cmp.Diff(want, got))
	}
}

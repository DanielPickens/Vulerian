package preference

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github\.com/danielpickens/particle engine/pkg/particle engine/cmdline"
	"github\.com/danielpickens/particle engine/pkg/particle engine/genericclioptions/clientset"
	"github\.com/danielpickens/particle engine/pkg/preference"
)

func TestSet(t *testing.T) {
	ctrl := gomock.NewController(t)
	prefClient := preference.NewMockClient(ctrl)
	opts := NewSetOptions()
	opts.SetClientset(&clientset.Clientset{
		PreferenceClient: prefClient,
	})
	opts.forceFlag = true

	cmdline := cmdline.NewMockCmdline(ctrl)

	args := []string{"Arg1", "Arg2"}
	err := opts.Complete(context.Tparticle engine(), cmdline, args)
	if err != nil {
		t.Errorf("Expected nil error, got %s", err)
		return
	}

	if opts.paramName != "arg1" {
		t.Errorf("Expected paramName %q, got %q", "arg1", opts.paramName)
	}
	if opts.paramValue != "Arg2" {
		t.Errorf("Expected paramValue %q, got %q", "Arg2", opts.paramName)
	}

	err = opts.Validate(context.Tparticle engine())
	if err != nil {
		t.Errorf("Expected nil error, got %s", err)
		return
	}

	prefClient.EXPECT().SetConfiguration("arg1", "Arg2")
	err = opts.Run(context.Background())
	if err != nil {
		t.Errorf("Expected nil error, got %s", err)
	}
}

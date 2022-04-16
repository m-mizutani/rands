package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/m-mizutani/rands"
	"github.com/urfave/cli/v2"
)

type config struct {
	Length      int
	Seed        int64
	CharSet     string
	AppendChars cli.StringSlice
	Output      string

	UseLowers  bool
	UseUppers  bool
	UseNumbers bool
	UseMarks   bool
}

func (x *config) baseChars() string {
	switch {
	case x.CharSet != "":
		return x.CharSet

	case x.UseLowers || x.UseUppers || x.UseNumbers || x.UseMarks:
		var chars string
		if x.UseLowers {
			chars += rands.LowerSet
		}
		if x.UseUppers {
			chars += rands.UpperSet
		}
		if x.UseNumbers {
			chars += rands.NumberSet
		}
		if x.UseMarks {
			chars += rands.MarkSet
		}
		return chars

	default:
		return rands.DefaultCharSet
	}
}

func main() {
	var cfg config
	app := &cli.App{
		Name:  "rands",
		Usage: "Random string generator",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "length",
				Aliases:     []string{"l"},
				Usage:       "Set length of random string",
				EnvVars:     []string{"RANDS_LENGTH"},
				Value:       12,
				Destination: &cfg.Length,
			},
			&cli.StringSliceFlag{
				Name:        "append",
				Aliases:     []string{"a"},
				Usage:       "Apppend charactor set for random string",
				EnvVars:     []string{"RANDS_APPEND"},
				Destination: &cfg.AppendChars,
			},
			&cli.Int64Flag{
				Name:        "seed",
				Aliases:     []string{"s"},
				Usage:       "Sen seed for random",
				EnvVars:     []string{"RANDS_SEED"},
				Destination: &cfg.Seed,
			},
			&cli.StringFlag{
				Name:        "output",
				Aliases:     []string{"o"},
				Usage:       "Set output file ('-' is stdout)",
				EnvVars:     []string{"RANDS_OUTPUT"},
				Destination: &cfg.Output,
				Value:       "-",
			},

			&cli.StringFlag{
				Name:        "chars",
				Aliases:     []string{"c"},
				Usage:       "Set charactor set for random string (This option is prioritized than --use-* options)",
				EnvVars:     []string{"RANDS_CHARS"},
				Destination: &cfg.CharSet,
			},

			// char set
			&cli.BoolFlag{
				Name:        "use-lowers",
				Usage:       "Use lower alphabet",
				EnvVars:     []string{"RANDS_USE_LOWER"},
				Destination: &cfg.UseLowers,
			},
			&cli.BoolFlag{
				Name:        "use-uppers",
				Usage:       "Use upper alphabet",
				EnvVars:     []string{"RANDS_USE_Upper"},
				Destination: &cfg.UseUppers,
			},
			&cli.BoolFlag{
				Name:        "use-numbers",
				Usage:       "Use numbers",
				EnvVars:     []string{"RANDS_USE_NUMBERS"},
				Destination: &cfg.UseNumbers,
			},
			&cli.BoolFlag{
				Name:        "use-marks",
				Usage:       "Use marks",
				EnvVars:     []string{"RANDS_USE_MARKS"},
				Destination: &cfg.UseMarks,
			},
		},
		Action: func(ctx *cli.Context) error {
			// setup
			chars := cfg.baseChars()
			for _, append := range cfg.AppendChars.Value() {
				chars += append
			}

			options := []rands.Option{
				rands.WithCharSet(chars),
			}
			if cfg.Seed > 0 {
				options = append(options, rands.WithSeed(cfg.Seed))
			}

			r := rands.New(options...)

			// output
			var out io.Writer
			if cfg.Output == "-" {
				out = os.Stdout
			} else {
				fd, err := os.Create(cfg.Output)
				if err != nil {
					return fmt.Errorf("failed to open: %w", err)
				}
				defer func() {
					if err := fd.Close(); err != nil {
						log.Println("failed to close:", err.Error())
					}
				}()
			}

			random := r.NewString(cfg.Length)
			if _, err := fmt.Fprintln(out, random); err != nil {
				return fmt.Errorf("failed to output: %w", err)
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("runtime error: %+v\n", err)
	}
}

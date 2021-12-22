package fluit

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/term"
)

const on_error_value int = 80
const arg_padding_length int = 2

func fmt_margin_wrap(input_text string, margin_length int, line_break_at int) (res string, err error) {
	if margin_length < 0 {
		err = fmt.Errorf("%werror Sprintfm(): negative margin_length; ", err)
		fmt.Println(err)
		return "", err
	}
	if line_break_at <= 0 {
		line_break_at = on_error_value
	}
	if margin_length > line_break_at {
		err = fmt.Errorf("%werror Sprintfmt(): margin_length larger than line_break, using terminal width as line_break; ", err)
		line_break_at, _, term_getsize_err := term.GetSize(1)
		if term_getsize_err != nil {
			err = fmt.Errorf("%werror term.GetSize(): %q; ", err, term_getsize_err)
			err = fmt.Errorf("%werror Sprintfmt(): can't get terminal size; ", err)
			return input_text + "\r\n", err
		}
		if margin_length > line_break_at {
			err = fmt.Errorf("%werror Sprintfmt(): margin_length larger than terminal width, set margin to 0; ", err)
			margin_length = 0
		}
	}
	margin := strings.Repeat(" ", margin_length)
	if line_break_at > len(input_text)+margin_length {
		return margin + input_text + "\r\n", err
	}
	max_text_length := strconv.Itoa(line_break_at - margin_length)
	if !strings.Contains(input_text, " ") {
		regex_p := regexp.MustCompile("(.{1," + max_text_length + "})")
		return regex_p.ReplaceAllString(input_text, margin+"$1\r\n"), err
	}
	regex_p := regexp.MustCompile("(.{1," + max_text_length + "})(?:\\s|$)")
	return regex_p.ReplaceAllString(input_text, margin+"$1\r\n"), err
}

func sprintf_usg_verbose(arg string, desc string, max_arg_length int, line_break_at int) (res string, err error) {
	if max_arg_length < 0 {
		err = fmt.Errorf("error sprintf_usg(): max_arg_length can't be negative.")
		fmt.Println(err)
		return "", err
	} else if max_arg_length == 0 {
		max_arg_length = 24
	}
	if line_break_at <= 0 {
		line_break_at = on_error_value
	}
	if max_arg_length > line_break_at {
		err = fmt.Errorf("error sprintf_usg(): max_arg_length can't be larger than line_break length.")
		fmt.Println(err)
		return "", err
	}
	arg_collumn_length := max_arg_length + (arg_padding_length * 2)
	formatted_desc, err := fmt_margin_wrap(desc, arg_collumn_length, line_break_at)
	arg_length := len(arg)
	{
		if arg_length > max_arg_length {
			formatted_arg, _ := fmt_margin_wrap(arg, arg_padding_length, line_break_at)
			return formatted_arg + formatted_desc, err
		}
	}
	arg = strings.Repeat(" ", arg_padding_length) + arg + strings.Repeat(" ", max_arg_length-arg_length+arg_padding_length)
	res = strings.Replace(formatted_desc, strings.Repeat(" ", arg_collumn_length), arg, 1)
	return res, err
}

func Sprintfm(input_text string, margin_length int, line_break_at int) (res string) {
	res, _ = fmt_margin_wrap(input_text, margin_length, line_break_at)
	return
}

func Sprintf_usg(arg string, desc string, max_arg_length int, line_break_at int) (res string) {
	res, _ = sprintf_usg_verbose(arg, desc, max_arg_length, line_break_at)
	return
}

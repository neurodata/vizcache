package main

import (
	"fmt"
	"strconv"

	"net/http"
	"regexp"
	"strings"

	"github.com/labstack/echo"
)

type NDStoreContext struct {
	echo.Context
	args *NDStoreArgs
}

type NDStoreArgs struct {
	X       [2]int
	Y       [2]int
	Z       [2]int
	Res     int
	Token   string
	Channel string
	Server  string
}

func checkArgsError(e error) error {
	return fmt.Errorf("failed to parse RESTful request argument: %s", e.Error())
}

func (c *NDStoreContext) GetVolume(argsString string) error {
	// process args
	err := c.processArgs(argsString)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	fmt.Println(c.args.Token)

	// get volume from cache

	return c.String(http.StatusOK, "hello")
}

func (c *NDStoreContext) processArgs(argsString string) error {
	re := regexp.MustCompile(`^\/?(?P<server>[\w.:]+)?\/(?P<token>[\w]+)\/(?P<channel>[\w]+)\/(?P<res>[\d]+)\/(?P<x>[\d,]+)\/(?P<y>[\d,]+)\/(?P<z>[\d,]+)\/?$`)

	arr := re.FindStringSubmatch(argsString) //, -1)
	if arr == nil || len(arr) < 8 {
		return fmt.Errorf("error: Failed to parse RESTful arguments.\n%s\nReceived: %s", NDStoreArgsErr.VOLUME, argsString)
	}

	// initialize an empty NDStoreArgs struct
	c.args = new(NDStoreArgs)

	var err error

	c.args.Server = arr[1]
	c.args.Token = arr[2]
	c.args.Channel = arr[3]

	c.args.Res, err = strconv.Atoi(arr[4])
	if err != nil {
		return checkArgsError(err)
	}

	c.args.X, err = processCoordinateArgs(arr[5])
	if err != nil {
		return checkArgsError(err)
	}
	c.args.Y, err = processCoordinateArgs(arr[6])
	if err != nil {
		return checkArgsError(err)
	}
	c.args.Z, err = processCoordinateArgs(arr[7])
	if err != nil {
		return checkArgsError(err)
	}

	fmt.Println(c.args)

	return nil
}

func processCoordinateArgs(args string) ([2]int, error) {
	var output [2]int
	var err error
	arr := strings.Split(args, ",")
	for i := 0; i < 2; i++ {
		output[i], err = strconv.Atoi(arr[i])
		if err != nil {
			return output, err
		}
	}
	return output, nil
}

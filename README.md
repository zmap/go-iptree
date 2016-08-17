GoLang IPTree
=============

[![Build Status](https://travis-ci.org/zmap/go0iptree.svg?branch=master)](https://travis-ci.org/zmap/go-iptree)

This is a golang based prefix tree for IP subnets

Install
=======

ZDNS can be used by including:

	import "github.com/zmap/go-iptree"


Usage
=====

Below is a simple example:

	t := iptree.New()
	t.Insert("0.0.0.0", 0)
	t.Insert("128.255.0.0/16", 1)
	t.Insert("128.255.134.0/24", 0)

	if val, err := t.Get("128.255.134.5"); err != nil {
		fmt.Println("Value is", val)
	}

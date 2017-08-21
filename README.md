GoLang IPTree
=============

[![Build Status](https://travis-ci.org/zmap/go-iptree.svg?branch=master)](https://travis-ci.org/zmap/go-iptree)

This is a golang based prefix tree for IP subnets

Install
=======

go-iptree can be used by including:

	import "github.com/zmap/go-iptree"


Usage
=====

Below is a simple example:

	t := iptree.New()
	t.AddByString("0.0.0.0", 0)
	t.AddByString("128.255.0.0/16", 1)
	t.AddByString("128.255.134.0/24", 0)

	if val, found, err := t.GetByString("128.255.134.5"); err == nil && found {
		fmt.Println("Value is ", val)
	}

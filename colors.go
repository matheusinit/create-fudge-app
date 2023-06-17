package main

const Reset = "\033[0m"
const Green = "\033[32m"
const Yellow = "\033[33m"
const Blue = "\033[34m"
const Purple = "\033[35m"
const Cyan = "\033[36m"
const Gray = "\033[37m"
const White = "\033[38m"
const Red = "\033[0;31m"

func colorRed(text string) string {
	return Red + text + Red
}

function getBrowserInstance(): typeof chrome {
	const browserInstance = chrome || (Browser as any)
	return browserInstance
}

const Browser = getBrowserInstance()

export default Browser

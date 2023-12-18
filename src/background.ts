import Browser from "./modules/browser"

Browser.action.setBadgeBackgroundColor({
	color: "#ae0000",
})

Browser.action.setBadgeTextColor({
	color: "#ffffff",
})

if (Browser.declarativeNetRequest.setExtensionActionOptions) {
	Browser.declarativeNetRequest.setExtensionActionOptions({
		displayActionCountAsBadgeText: true,
	})
}

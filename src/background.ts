import Browser from "./modules/browser"

if (Browser.action.setBadgeBackgroundColor) {
	Browser.action.setBadgeBackgroundColor({
		color: "#ae0000",
	})
}

if (Browser.action.setBadgeTextColor) {
	Browser.action.setBadgeTextColor({
		color: "#ffffff",
	})
}

if (Browser.declarativeNetRequest.setExtensionActionOptions) {
	Browser.declarativeNetRequest.setExtensionActionOptions({
		displayActionCountAsBadgeText: true,
	})
}

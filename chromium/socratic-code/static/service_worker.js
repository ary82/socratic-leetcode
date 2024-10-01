let state = null;

chrome.runtime.onMessage.addListener((message, _sender, sendResponse) => {
	if (message.type === 'update-data') {
		state = message.data;
	} else if (message.type === 'get-data') {
		sendResponse(state);
	}
	return true;
});

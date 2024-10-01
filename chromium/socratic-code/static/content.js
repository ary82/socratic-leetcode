let problemText;
let codeText;
let langText;
let questionText;
let codeHtml;

function sendData() {
	chrome.runtime.sendMessage({
		type: 'update-data',
		data: {
			problemText,
			codeText,
			langText,
			questionText,
			codeHtml
		}
	});
}

const interval = setInterval(() => {
	let problemEle = document.querySelector('[data-track-load="description_content"]');
	let codeEle = document.querySelector('#editor .view-lines');
	let langEle = document.querySelector('#editor [id^="headlessui"]');
	let questionEle = document.querySelector('[data-layout-path="/ts0/t0"] a');

	if (problemEle && codeEle && langEle && questionEle) {
		problemText = problemEle.innerText;
		codeText = codeEle.innerText;
		langText = langEle.innerText;
		questionText = questionEle.href;
		codeHtml = codeEle.innerHTML;
		sendData();
	}
}, 500);

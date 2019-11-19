var xhr = new XMLHttpRequest()
var url = 'http://localhost:8080/'

const handler = () => {
    // コンソールに出力
    console.log(xhr.responseText)
    $("#status").append(xhr.responseText);
}

const getRequest = () => {
    xhr.open('GET', url)
    xhr.onloadend = handler
    xhr.send()
}

document.addEventListener('DOMContentLoaded', () => {
    getRequest()
})

function showMessage(msg) {
    document.getElementById('msg').innerText = msg;
}

// 工具函数：对象转 x-www-form-urlencoded 字符串
function encodeForm(data) {
    return Object.entries(data)
        .map(([k, v]) => encodeURIComponent(k) + '=' + encodeURIComponent(v))
        .join('&');
}

function register() {
    const username = document.getElementById('regUsername').value;
    const password = document.getElementById('regPassword').value;

    fetch('/register', {
        method: 'POST',
        headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
        body: encodeForm({ username, password }),
    })
        .then(res => res.text())
        .then(showMessage);
}

function requestResetToken() {
    const username = document.getElementById('resetUser1').value;

    fetch('/resetrequest', {
        method: 'POST',
        headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
        body: encodeForm({ username }),
    })
        .then(res => res.text())
        .then(showMessage);
}
function login() {
    const username = document.getElementById('loginUsername').value;
    const password = document.getElementById('loginPassword').value;

    fetch('/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
        body: encodeForm({ username, password }),
    })
        .then(res => res.text())
        .then(text => {
            showMessage(text);
            if (text === '登录成功') {
                window.location.href = '/';  // 或你想跳转的页面
            }
        });
}
function resetPassword() {
    const username = document.getElementById('resetUser2').value;
    const token = document.getElementById('resetToken').value;
    const newPassword = document.getElementById('newPassword').value;

    fetch('/resetconfirm', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ username, token, newPassword }),  // 发送 JSON 字符串
    })
        .then(res => res.text())
        .then(showMessage);
}

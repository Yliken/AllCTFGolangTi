document.getElementById('login-form').addEventListener('submit', function (e) {
    e.preventDefault(); // 防止表单默认提交

    const username = document.getElementById('username').value.trim();
    const studentnum = document.getElementById('studentnum').value.trim();

    fetch('/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ username, studentnum })
    })
        .then(res => res.json())
        .then(data => {
            const msg = document.getElementById('login-message');
            if (data.success) {
                msg.style.color = 'green';
                msg.innerText = '登录成功！正在跳转...';
                setTimeout(() => {
                    window.location.href = '/';
                }, 1000);
            } else {
                msg.style.color = 'red';
                msg.innerText = data.message || '用户名或学号错误';
            }
        })
        .catch(() => {
            document.getElementById('login-message').innerText = '请求失败，请稍后再试';
        });
});

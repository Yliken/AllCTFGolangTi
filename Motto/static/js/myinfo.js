document.getElementById('changeNicknameBtn').addEventListener('click', function() {
    const newNickname = document.getElementById('newNickname').value.trim();
    if (!newNickname) {
        alert('请输入新的昵称');
        return;
    }

    fetch('/changeNickName', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ nickname: newNickname })
    })
        .then(response => {
            if (response.ok) {
                document.getElementById('nickname').textContent = newNickname;
                alert('昵称修改成功');
            } else {
                alert('昵称修改失败');
            }
        })
        .catch(error => {
            alert('请求出错: ' + error.message);
        });
});

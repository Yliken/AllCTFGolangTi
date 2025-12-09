document.addEventListener("DOMContentLoaded", () => {
    const form = document.getElementById("registerForm");
    const nicknameInput = document.getElementById("nickname");
    const usernameInput = document.getElementById("username");
    const passwordInput = document.getElementById("password");

    const nicknameError = document.getElementById("nicknameError");
    const usernameError = document.getElementById("usernameError");
    const passwordError = document.getElementById("passwordError");

    form.addEventListener("submit", async (e) => {
        e.preventDefault(); // 阻止默认提交

        // 清空错误信息
        nicknameError.textContent = "";
        usernameError.textContent = "";
        passwordError.textContent = "";

        let valid = true;

        if (nicknameInput.value.trim() === "") {
            nicknameError.textContent = "请输入昵称";
            valid = false;
        }
        if (usernameInput.value.trim() === "") {
            usernameError.textContent = "请输入账号";
            valid = false;
        }
        if (passwordInput.value.trim() === "") {
            passwordError.textContent = "请输入密码";
            valid = false;
        } else if (passwordInput.value.length < 6) {
            passwordError.textContent = "密码至少6位";
            valid = false;
        }

        if (!valid) return;

        // 构造 application/x-www-form-urlencoded 格式的请求体
        const formData = new URLSearchParams();
        formData.append('nickname', nicknameInput.value.trim());
        formData.append('username', usernameInput.value.trim());
        formData.append('password', passwordInput.value);

        try {
            const response = await fetch('/register', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8'
                },
                body: formData.toString()
            });

            if (!response.ok) {
                throw new Error(`服务器返回错误: ${response.status}`);
            }

            const result = await response.json();

            if (result.success) {
                alert('注册成功！');
                form.reset();
                window.location.href = '/login';
            } else {
                alert(`注册失败: ${result.message || '未知错误'}`);
            }
        } catch (error) {
            alert(`请求失败: ${error.message}`);
        }
    });
});

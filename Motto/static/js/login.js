document.addEventListener("DOMContentLoaded", () => {
    const form = document.getElementById("loginForm");
    const usernameInput = document.getElementById("username");
    const passwordInput = document.getElementById("password");
    const usernameError = document.getElementById("usernameError");
    const passwordError = document.getElementById("passwordError");

    form.addEventListener("submit", async (e) => {
        e.preventDefault();

        // 清空之前的错误信息
        usernameError.textContent = "";
        passwordError.textContent = "";

        let valid = true;

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

        try {
            const response = await fetch('/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    username: usernameInput.value.trim(),
                    password: passwordInput.value
                })
            });

            if (!response.ok) {
                // HTTP错误状态处理
                throw new Error(`服务器返回错误: ${response.status}`);
            }

            const result = await response.json();

            // 假设后端返回格式 { success: true/false, message: "提示信息" }
            if (result.success) {
                alert('登录成功！');
                form.reset();
                // 登录成功后可以跳转页面，比如：
                window.location.href = '/';
            } else {
                alert(`登录失败: ${result.message || '未知错误'}`);
            }
        } catch (error) {
            alert(`请求失败: ${error.message}`);
        }
    });
});

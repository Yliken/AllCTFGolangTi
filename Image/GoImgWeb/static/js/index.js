document.addEventListener("DOMContentLoaded", () => {
    const uploadBtn = document.getElementById("uploadBtn");
    const uploadModal = document.getElementById("uploadModal");
    const closeModal = document.getElementById("closeModal");
    const fileInput = document.getElementById("fileInput");
    const submitUpload = document.getElementById("submitUpload");
    const uploadStatus = document.getElementById("uploadStatus");

    uploadBtn.addEventListener("click", () => {
        uploadModal.style.display = "block";
        uploadStatus.textContent = "";
        fileInput.value = "";
    });

    closeModal.addEventListener("click", () => {
        uploadModal.style.display = "none";
    });

    window.addEventListener("click", (e) => {
        if (e.target === uploadModal) {
            uploadModal.style.display = "none";
        }
    });

    submitUpload.addEventListener("click", () => {
        const file = fileInput.files[0];
        if (!file) {
            uploadStatus.style.color = "red";
            uploadStatus.textContent = "请选择一个图片文件！";
            return;
        }

        // 构造 FormData 对象上传文件
        const formData = new FormData();
        formData.append("file", file);

        uploadStatus.style.color = "black";
        uploadStatus.textContent = "上传中...";

        fetch("/upload", {
            method: "POST",
            body: formData,
        })
            .then(response => response.json())
            .then(data => {
                if (data.success === true) {
                    uploadStatus.style.color = "green";
                    uploadStatus.textContent = "上传成功！";

                    // 上传成功后关闭弹窗并刷新页面或重新加载图片列表
                    setTimeout(() => {
                        uploadModal.style.display = "none";
                        uploadStatus.textContent = "";
                        location.reload(); // 重新加载页面展示新图片
                    }, 1500);
                } else {
                    alert("请先登录");
                    window.location.href = "/login";
                }
            })
            .catch(err => {
                console.error(err);
                uploadStatus.style.color = "red";
                uploadStatus.textContent = "上传异常，请重试！";
            });
    });
});

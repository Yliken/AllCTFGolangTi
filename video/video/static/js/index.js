// 全局变量
let isPlaying = false;
let isMuted = true; // 默认静音
let isLoading = false;

// DOM元素
const videoPlayer = document.getElementById('videoPlayer');
const playButton = document.getElementById('playButton');
const loadingOverlay = document.getElementById('loadingOverlay');
const errorOverlay = document.getElementById('errorOverlay');
const errorText = document.getElementById('errorText');
const nextBtn = document.getElementById('nextBtn');
const muteBtn = document.getElementById('muteBtn');

// 获取视频的函数
async function fetchVideo() {
    if (isLoading) return;

    setLoading(true);
    hideError();

    try {
        const response = await fetch('/getvideo');

        if (!response.ok) {
            throw new Error(`HTTP ${response.status}: ${response.statusText}`);
        }

        const data = await response.json();
        const videoUrl = data.url || data.videoUrl || data.video || data;

        if (!videoUrl) {
            throw new Error('响应中没有找到视频URL');
        }

        // 设置视频源
        videoPlayer.src = videoUrl;

    } catch (error) {
        console.error('获取视频失败:', error);
        showError('无法获取视频: ' + error.message);

        // 使用示例视频作为fallback
        videoPlayer.src = 'https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4';
    }
}

// 设置加载状态
function setLoading(loading) {
    isLoading = loading;
    nextBtn.disabled = loading;

    if (loading) {
        loadingOverlay.classList.remove('hidden');
    } else {
        loadingOverlay.classList.add('hidden');
    }
}

// 显示错误
function showError(message) {
    errorText.textContent = message;
    errorOverlay.classList.remove('hidden');
    setLoading(false);
}

// 隐藏错误
function hideError() {
    errorOverlay.classList.add('hidden');
}

// 播放/暂停切换
function togglePlay() {
    if (!videoPlayer.src) return;

    if (isPlaying) {
        videoPlayer.pause();
    } else {
        videoPlayer.play().catch(error => {
            console.error('播放失败:', error);
            showError('视频播放失败: ' + error.message);
        });
    }
}

// 静音切换
function toggleMute() {
    isMuted = !isMuted;
    videoPlayer.muted = isMuted;
    updateMuteButton();
}

// 更新静音按钮显示
function updateMuteButton() {
    const volumeOn = muteBtn.querySelector('.volume-on');
    const volumeOff = muteBtn.querySelector('.volume-off');

    if (isMuted) {
        volumeOn.classList.add('hidden');
        volumeOff.classList.remove('hidden');
    } else {
        volumeOn.classList.remove('hidden');
        volumeOff.classList.add('hidden');
    }
}

// 切换到下一个视频
function nextVideo() {
    if (isLoading) return;

    // 暂停当前视频
    videoPlayer.pause();
    setPlaying(false);

    // 获取新视频
    fetchVideo();
}

// 设置播放状态
function setPlaying(playing) {
    isPlaying = playing;
    updatePlayButton();
}

// 更新播放按钮显示
function updatePlayButton() {
    const playIcon = playButton.querySelector('.play-icon');
    const pauseIcon = playButton.querySelector('.pause-icon');

    if (isPlaying) {
        playIcon.classList.add('hidden');
        pauseIcon.classList.remove('hidden');
        // 隐藏播放按钮整体
        playButton.classList.add('hidden');
    } else {
        playIcon.classList.remove('hidden');
        pauseIcon.classList.add('hidden');
        // 显示播放按钮整体
        playButton.classList.remove('hidden');
    }
}

// 视频事件监听器
videoPlayer.addEventListener('loadstart', () => {
    console.log('开始加载视频');
});

videoPlayer.addEventListener('loadeddata', () => {
    console.log('视频数据加载完成');
    setLoading(false);
    hideError();

    // 自动播放
    videoPlayer.play().catch(error => {
        console.error('自动播放失败:', error);
        // 某些浏览器需要用户交互才能播放
    });
});

videoPlayer.addEventListener('play', () => {
    console.log('视频开始播放');
    setPlaying(true);
});

videoPlayer.addEventListener('pause', () => {
    console.log('视频暂停');
    setPlaying(false);
});

videoPlayer.addEventListener('ended', () => {
    console.log('视频播放结束');
    setPlaying(false);

    // 自动获取下一个视频
    setTimeout(() => {
        nextVideo();
    }, 500);
});

videoPlayer.addEventListener('error', (e) => {
    console.error('视频播放错误:', e);
    const error = videoPlayer.error;
    let errorMessage = '视频播放出错';

    if (error) {
        switch (error.code) {
            case error.MEDIA_ERR_ABORTED:
                errorMessage = '视频播放被中止';
                break;
            case error.MEDIA_ERR_NETWORK:
                errorMessage = '网络错误导致视频下载失败';
                break;
            case error.MEDIA_ERR_DECODE:
                errorMessage = '视频解码失败';
                break;
            case error.MEDIA_ERR_SRC_NOT_SUPPORTED:
                errorMessage = '视频格式不支持';
                break;
            default:
                errorMessage = '未知的视频播放错误';
        }
    }

    showError(errorMessage);
});

// 键盘事件监听
document.addEventListener('keydown', (e) => {
    switch (e.code) {
        case 'Space':
            e.preventDefault();
            togglePlay();
            break;
        case 'ArrowUp':
        case 'ArrowDown':
            e.preventDefault();
            nextVideo();
            break;
        case 'KeyM':
            e.preventDefault();
            toggleMute();
            break;
    }
});

// 触摸事件处理（移动端）
let touchStartY = 0;
let touchEndY = 0;

document.addEventListener('touchstart', (e) => {
    touchStartY = e.changedTouches[0].screenY;
});

document.addEventListener('touchend', (e) => {
    touchEndY = e.changedTouches[0].screenY;
    handleSwipe();
});

function handleSwipe() {
    const swipeThreshold = 50;
    const diff = touchStartY - touchEndY;

    if (Math.abs(diff) > swipeThreshold) {
        // 上滑或下滑切换视频
        nextVideo();
    }
}

// 页面初始化
document.addEventListener('DOMContentLoaded', () => {
    console.log('页面加载完成，初始化视频播放器');

    // 设置初始状态
    videoPlayer.muted = isMuted;
    updateMuteButton();
    updatePlayButton();

    // 获取第一个视频
    fetchVideo();
});

// 页面可见性变化处理
document.addEventListener('visibilitychange', () => {
    if (document.hidden) {
        // 页面隐藏时暂停视频
        if (isPlaying) {
            videoPlayer.pause();
        }
    }
});

// 窗口失焦时暂停视频
window.addEventListener('blur', () => {
    if (isPlaying) {
        videoPlayer.pause();
    }
});

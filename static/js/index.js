// 从 DOM 中获取数据
const usageSection = document.querySelector('.usage-section');
const fileDetailSection = document.querySelector('.file-detail');

const currentSize = usageSection.getAttribute('data-current-size');
const maxSize = usageSection.getAttribute('data-max-size');

const fileDetails = {
    docCount: fileDetailSection.getAttribute('data-doc-count'),
    imgCount: fileDetailSection.getAttribute('data-img-count'),
    videoCount: fileDetailSection.getAttribute('data-video-count'),
    musicCount: fileDetailSection.getAttribute('data-music-count'),
    otherCount: fileDetailSection.getAttribute('data-other-count')
};

// 容量图表
const ctxCapacity = document.getElementById('capacityChart').getContext('2d');
const capacityChart = new Chart(ctxCapacity, {
    type: 'doughnut',
    data: {
        labels: ['已使用容量', '剩余容量'],
        datasets: [{
            data: [currentSize, maxSize - currentSize],
            backgroundColor: ['#4caf50', '#f44336']
        }]
    },
    options: {
        responsive: true,
        plugins: {
            legend: {
                position: 'top',
            },
        }
    }
});

// 文件详情图表
const ctxFileDetail = document.getElementById('fileDetailChart').getContext('2d');
const fileDetailChart = new Chart(ctxFileDetail, {
    type: 'bar',
    data: {
        labels: ['文本', '图像', '视频', '音乐', '其他'],
        datasets: [{
            label: '文件数量',
            data: [fileDetails.docCount, fileDetails.imgCount, fileDetails.videoCount, fileDetails.musicCount, fileDetails.otherCount],
            backgroundColor: [
                '#2196f3',
                '#ff9800',
                '#f44336',
                '#4caf50',
                '#9c27b0'
            ]
        }]
    },
    options: {
        responsive: true,
        plugins: {
            legend: {
                display: false
            }
        },
        scales: {
            y: {
                beginAtZero: true
            }
        }
    }
});
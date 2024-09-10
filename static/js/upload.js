document.addEventListener('DOMContentLoaded', function() {
    const form = document.getElementById('uploadForm');
    const uploadButton = document.getElementById('uploadButton');
    const statusDiv = document.getElementById('status');

    uploadButton.addEventListener('click', function() {
        const formData = new FormData(form);
        const fid = formData.get('fid');

        if (!formData.get('file')) {
            statusDiv.textContent = '请先选择一个文件';
            return;
        }

        fetch('/cloud/uploadFile', {
            method: 'POST',
            headers: {
                'fid': fid
            },
            body: formData
        })
            .then(response => response.json())
            .then(data => {
                if (data.error){
                    statusDiv.textContent='文件上传失败，原因：'+data.error;
                }else{
                    statusDiv.textContent = '文件上传成功';
                }
            })
            .catch(error => {
                statusDiv.textContent = '因网络原因文件上传失败';
            });
    });
});
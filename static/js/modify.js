document.getElementById('updateUserForm').addEventListener('submit', function (event) {
    event.preventDefault(); // 阻止默认提交行为

    const currentPassword = document.getElementById('currentPassword').value;
    const newPassword =document.getElementById('newPassword').value;
    const confirmPassword = document.getElementById('confirmPassword').value;
    const nickname = document.getElementById('nickname').value.trim();
    const avatar = document.getElementById('avatar').files[0];
    const errorMessage = document.getElementById('error-message');

    // 如果输入新密码需验证原密码输入以及确认密码是否匹配
    if (newPassword.length>0){
        if (currentPassword.length===0){
            errorMessage.textContent = '请输入原密码！';
            return;
        }
        if (newPassword!==confirmPassword){
            errorMessage.textContent = '新密码和确认密码不匹配！';
            return;
        }
    }

    // 验证昵称是否为 1-10 位
    if (nickname.length > 10) {
        errorMessage.textContent = '昵称必须为 1 到 10 位！';
        return;
    }

    // 验证头像文件格式
    if (avatar) {
        const allowedTypes = ['image/jpeg', 'image/png', 'image/gif'];
        if (!allowedTypes.includes(avatar.type)) {
            errorMessage.textContent = '头像文件格式不正确，请上传 JPEG、PNG 或 GIF 图片！';
            return;
        }
    }

    // 如果所有验证都通过，清除错误信息并提交表单
    errorMessage.textContent = '';

    // 提交表单
    this.submit();
});
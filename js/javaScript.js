	
function validation(){
	var username=document.getElementById('username').value;
	var password=document.getElementById('password').value;
	console.log(username);

	if(password.length<4||password.length>12)
	{
		alert("Mật khẩu phải có từ 4 đến 12 ký tự");
		return false;
	}
	else
	{

		var str=username.slice(0,1);
		if(username.slice(0,1)=="_"||username.slice(0,1)=="@"||str.match(/[0-9]/g) !=null)
		{
			alert("username should not start with _, @ and number");
			return false;
		}
		else
		{
			if(username=="admin" && password=="admin")
			{
				document.location = "index.html";
				return false;
			}
			else
			{
				alert("Sai mật khẩu hoặc tên đăng nhập");
				return false;
			}
		}

	}
}
function signUpValidation()
{
	var username = document.getElementById('username').value;
	var password = document.getElementById('password').value;
	var confirm = document.getElementById('confirmPassword').value;
	if(username == "admin")
	{
		alert("Tên đăng nhập đã có người sử dụng.");
		return false;
	}
	else
	{
		if(password.length<4||password.length>12)
		{
			alert("Mật khẩu phải có từ 4 đến 12 ký tự");
			return false;
		}
		if(password != confirm)
		{
			alert("Mật khẩu nhập lại không chính xác.");
			return false;
		}

	}
	return true;

}
var JdeConfiguration = function() {
   var visibleMenu = function(userRole) {
      if(userRole == "Admin") {
         model.MainMenus.push(new MenuItem("usermanagement", "/user/default", "User Management", "usernamagement", []));
      }
   }

   function validationByUsername(param, tag) {
      var url = "/user/cekbyusername";
      var success;
     ajaxPost(url, param, function(data){
         success = data.Success;
         if(!success) {
            $("#"+tag).focus();
            model.User1.Username("");
            swal("Username already use", "", 'error');
         }
      });
   }

   function uploadImageProfile(tag, name, url, role) {
      var imgdata = new FormData();
      imgdata.append("userfile", tag[0].files[0]);
      imgdata.append("username", name);
      if (tag[0].files[0] != undefined){
         imgdata.append("filetype", (tag[0].files[0].name).split('.').pop());
      }else{
         imgdata.append("filetype", '');
      }

      var request = new XMLHttpRequest();
      request.open("POST", url);
      request.send(imgdata);
      request.onreadystatechange = function(res) {
         if(request.readyState == 4 && request.status == 200) {
            var data = JSON.parse(request.responseText);

            var urlImage = "/static" + "/" + data.pathPhoto;

            if(role != "Admin")
            {
               $("#user-icon-photo")
                  .attr('src', urlImage)
                  .width(64)
                  .height(64);
            }
         }
      }
   }

   var init = function() {
   };

   $(init);

   return {
      visibleMenu : visibleMenu,
      validationByUsername : validationByUsername,
      uploadImageProfile : uploadImageProfile
   };
}();
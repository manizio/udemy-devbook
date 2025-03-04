$("#unfollow").on("click", unfollow)
$("#follow").on("click", follow)
$("#edit-user").on("submit", edit)
$("#update-password").on("submit", updatePassword)
$("#delete-user").on("click", deleteUser)

function unfollow() {
    const userID = $(this).data("user-id")
    $(this).prop('disabled', true)

    $.ajax({
        url: `/users/${userID}/unfollow`,
        method: "POST"
    }).done(function() {
        window.location = `/users/${userID}`
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao parar de seguir o usuário", "error")
        $('#unfollow').prop('disabled', false)
    })
}

function follow() {
    const userID = $(this).data('user-id')
    $(this).prop('disabled', true)

    $.ajax({
        url: `/users/${userID}/follow`,
        method: "POST"
    }).done(function() {
        window.location = `/users/${userID}`
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao seguir o usuário", "error")
        $('#follow').prop('disabled', false)
    })

}

function edit(event) {
    event.preventDefault()

    $.ajax({
        url: "/edit-user",
        method: "PUT",
        data: {
            name: $('#name').val(),
            email: $('#email').val(),
            nick: $('#nick').val()
        }
    }).done(function() {
        Swal.fire("Sucesso!", "Usuário atualizado com sucesso", "success").then(
            function(){
                window.location = "/profile"
            }
        )
    }).fail(function (){
        Swal.fire("Ops...", "Erro ao atualizar o usuário", "error")
    })
}

function updatePassword(event) {
    event.preventDefault();

    if ($('#new-password').val() != $('#confirm-password').val()) {
        Swal.fire("Ops...", "As senhas não concidem", "warning")
        return
    }

    $.ajax({
        url: "/update-password",
        method: "POST",
        data: {
            current: $('#current-password').val(),
            new: $('#new-password').val()
        }
    }).done(function() {
        Swal.fire("Sucesso!", "A senha foi atualizada com sucesso", "success").then(
            function(){
                window.location = "/profile"
            }
        )
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao atualizar a senha!", "error")
    })
}

function deleteUser(){

    Swal.fire({
        title: "Atenção!",
        text: "Tem certeza que deseja apagar a sua conta? Essa é uma ação irreversível",
        showCancelButton: true,
        cancelButtonText: "Cancel",
        icon: "warning"
    }).then(function(confirm){
        if (confirm.value) {
            $.ajax({
                url: "/delete-user",
                method: "DELETE"
            }).done(function() {
                Swal.fire("Sucesso!", "Seu usuário foi excluído com sucesso!", "success").then(
                    function(){
                        window.location = "/logout"
                    }
                )
            }).fail(function() {
                Swal.fire("Ops...", "Ocorreu um erro ao excluir o seu usuário!", "error")
            })
        }
    })
}

$('#formulario-cadastro').on('submit', createUser)

function createUser(e) {
    e.preventDefault();

    if ($('#password').val() != $("#confirm-password").val()) {
        alert("As senhas não coincidem");
        return;
    }

    $.ajax({
        url: "/users",
        method: "POST",
        data: {
            name: $('#name').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
            password: $('#password').val(),
        }
    }).done(function() {
        alert("Usuário cadastrado com sucesso")
    }).fail(function() {
        alert("Erro ao cadastrar usuário")
    })
}

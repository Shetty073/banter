defmodule UserServiceWeb.UserController do
  use UserServiceWeb, :controller

  def users(conn, _params) do
    users = [
      %{id: 1, name: "Ashish", email: "shetty073@gmail.com"},
      %{id: 2, name: "Ria", email: "riashetty96@gmail.com"}
    ]

    json(conn, %{users: users})

  end

end

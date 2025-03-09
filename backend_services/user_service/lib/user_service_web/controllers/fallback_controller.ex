defmodule UserServiceWeb.FallbackController do
  use UserServiceWeb, :controller

  def not_found(conn, _params) do
    conn
    |> put_status(:not_found)
    |> json(%{
      error: "Route not found",
      message: "The route you are trying to access does not exist.",
      path: conn.request_path
    })
  end
end

defmodule NotificationServiceWeb.Router do
  use NotificationServiceWeb, :router

  pipeline :api do
    plug :accepts, ["json"]
  end

  scope "/api", NotificationServiceWeb do
    pipe_through :api
  end

  # Enable Swoosh mailbox preview in development
  if Application.compile_env(:notification_service, :dev_routes) do

    scope "/dev" do
      pipe_through [:fetch_session, :protect_from_forgery]

      forward "/mailbox", Plug.Swoosh.MailboxPreview
    end
  end
end

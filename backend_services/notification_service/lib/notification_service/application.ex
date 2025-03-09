defmodule NotificationService.Application do
  # See https://hexdocs.pm/elixir/Application.html
  # for more information on OTP Applications
  @moduledoc false

  use Application

  @impl true
  def start(_type, _args) do
    children = [
      NotificationServiceWeb.Telemetry,
      NotificationService.Repo,
      {DNSCluster, query: Application.get_env(:notification_service, :dns_cluster_query) || :ignore},
      {Phoenix.PubSub, name: NotificationService.PubSub},
      # Start the Finch HTTP client for sending emails
      {Finch, name: NotificationService.Finch},
      # Start a worker by calling: NotificationService.Worker.start_link(arg)
      # {NotificationService.Worker, arg},
      # Start to serve requests, typically the last entry
      NotificationServiceWeb.Endpoint
    ]

    # See https://hexdocs.pm/elixir/Supervisor.html
    # for other strategies and supported options
    opts = [strategy: :one_for_one, name: NotificationService.Supervisor]
    Supervisor.start_link(children, opts)
  end

  # Tell Phoenix to update the endpoint configuration
  # whenever the application is updated.
  @impl true
  def config_change(changed, _new, removed) do
    NotificationServiceWeb.Endpoint.config_change(changed, removed)
    :ok
  end
end

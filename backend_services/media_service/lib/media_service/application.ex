defmodule MediaService.Application do
  # See https://hexdocs.pm/elixir/Application.html
  # for more information on OTP Applications
  @moduledoc false

  use Application

  @impl true
  def start(_type, _args) do
    children = [
      MediaServiceWeb.Telemetry,
      MediaService.Repo,
      {DNSCluster, query: Application.get_env(:media_service, :dns_cluster_query) || :ignore},
      {Phoenix.PubSub, name: MediaService.PubSub},
      # Start the Finch HTTP client for sending emails
      {Finch, name: MediaService.Finch},
      # Start a worker by calling: MediaService.Worker.start_link(arg)
      # {MediaService.Worker, arg},
      # Start to serve requests, typically the last entry
      MediaServiceWeb.Endpoint
    ]

    # See https://hexdocs.pm/elixir/Supervisor.html
    # for other strategies and supported options
    opts = [strategy: :one_for_one, name: MediaService.Supervisor]
    Supervisor.start_link(children, opts)
  end

  # Tell Phoenix to update the endpoint configuration
  # whenever the application is updated.
  @impl true
  def config_change(changed, _new, removed) do
    MediaServiceWeb.Endpoint.config_change(changed, removed)
    :ok
  end
end

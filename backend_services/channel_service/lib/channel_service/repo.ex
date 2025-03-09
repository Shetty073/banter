defmodule ChannelService.Repo do
  use Ecto.Repo,
    otp_app: :channel_service,
    adapter: Ecto.Adapters.Postgres
end

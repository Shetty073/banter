defmodule PresenceService.Repo do
  use Ecto.Repo,
    otp_app: :presence_service,
    adapter: Ecto.Adapters.Postgres
end

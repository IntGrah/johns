FROM astral/uv:python3.12-bookworm-slim

WORKDIR /app

COPY pyproject.toml uv.lock ./

RUN uv sync --frozen --no-cache

COPY johns.py .

CMD ["uv", "run", "python", "-m", "johns"]

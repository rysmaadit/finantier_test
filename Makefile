setup:
	cd auth_service && cp .env.example .env && cd .. && \
	cd encryption_service && cp .env.example .env && cd .. && \
	cd stock_service && cp .env.example .env

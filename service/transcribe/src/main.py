# Transcribe service driver
from fastapi import FastAPI
from src.api.v1.routers.transcribe import router as transcribe_router

app = FastAPI()
app.include_router(transcribe_router)

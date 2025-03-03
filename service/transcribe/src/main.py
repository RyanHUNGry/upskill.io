# Transcribe service driver
from fastapi import FastAPI
from src.api.v1.routers.transcribe import router as transcribe_router

app = FastAPI()
app.include_router(transcribe_router)

@app.get("/status")
async def get_status():
    return {"status": "ok"}

from fastapi import APIRouter, UploadFile

from src.model.transcribe import Transcribe

router = APIRouter(
    prefix="/transcribe",
)

@router.post("/")
async def transcribe_file(file: UploadFile):
    raw_audio = await file.read()
    await file.close()

    transcribe = Transcribe("tiny.en")
    audio = transcribe.transform_audio(raw_audio)
    transcribed_audio = transcribe.transcribe(audio)

    return {"name": file.filename, "response_text": transcribed_audio}

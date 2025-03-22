from fastapi import APIRouter, UploadFile, Depends

from typing import Annotated

from src.model.transcribe import Transcribe

router = APIRouter(
    prefix="/transcribe",
)

transcribe_annotation = Annotated[Transcribe, Depends(Transcribe)]

@router.post("/")
async def transcribe_file(file: UploadFile, model: transcribe_annotation):
    raw_audio = file.file
    audio = model.transform_audio(raw_audio)
    transcribed_audio = model.transcribe(audio)

    return {"response_text": transcribed_audio}

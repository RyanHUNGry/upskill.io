import whisper
import numpy as np
import soundfile as sf

import io
from tempfile import SpooledTemporaryFile

class Transcribe:
    def __init__(self, model: str = "tiny.en") -> None:
        self.model = whisper.load_model(model)

    def transcribe(self, audio: np.ndarray) -> str:
        result = self.model.transcribe(audio, fp16=False)
        return result["text"]
    
    def transform_audio(self, audio: SpooledTemporaryFile) -> np.ndarray:
        audio_stream = io.BytesIO(audio.read())
        audio.close()
        audio, _ = sf.read(audio_stream, dtype="float32")
        return audio

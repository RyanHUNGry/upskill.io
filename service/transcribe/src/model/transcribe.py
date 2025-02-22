import whisper
import numpy as np
import soundfile as sf

import io

class Transcribe:
    def __init__(self, model: str) -> None:
        self.model = whisper.load_model(model)

    def transcribe(self, audio: bytes) -> str:
        result = self.model.transcribe(audio, fp16=False)
        return result["text"]
    
    def transform_audio(self, audio: bytes) -> np.ndarray:
        audio_stream = io.BytesIO(audio)
        audio, _ = sf.read(audio_stream, dtype="float32")
        return audio

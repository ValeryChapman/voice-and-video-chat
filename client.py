#!/usr/bin/env python
# -*- coding: utf-8 -*-
from   websocket import create_connection
import threading
import pyaudio
import cv2
import pickle


class Client:
    
    def __init__(self):
        
        ip   = input("Server IP: ")
        room = input("Room ID: ")
        
        self.voice_ws = create_connection(f"ws://{ip}/ws/voice/{room}")
        print("Voice server —> OK")
        
        self.video_ws = create_connection(f"ws://{ip}/ws/video/{room}")
        print("Video server —> OK")
        
        chunk_size   = 1024
        audio_format = pyaudio.paInt16
        channels     = 1
        rate         = 20000

        # initialise microphone recording
        self.p                      = pyaudio.PyAudio()
        self.voice_playing_stream   = self.p.open(format=audio_format, channels=channels, rate=rate, output=True, frames_per_buffer=chunk_size)
        self.voice_recording_stream = self.p.open(format=audio_format, channels=channels, rate=rate, input=True,  frames_per_buffer=chunk_size)
        print("Sound —> OK")
        
        # initialise video recording
        self.video_playing_stream = cv2.VideoCapture(0)
        print(f"Camera —> OK ({self.video_playing_stream.get(cv2.CAP_PROP_FPS)} FPS)")

        # start threads
        voice_receive_thread = threading.Thread(target=self.receive_voice_server_data).start()
        video_receive_thread = threading.Thread(target=self.receive_video_server_data).start()
        
        voice_send_thread = threading.Thread(target=self.send_voice_data_to_server).start()
        video_send_thread = threading.Thread(target=self.send_video_data_to_server).start()

    # get the voice data from the server
    def receive_voice_server_data(self):
        while True:
            data = self.voice_ws.recv()
            self.voice_playing_stream.write(data)
    
    # get the video data from the server
    def receive_video_server_data(self):
        while True:
            data = self.video_ws.recv()
            if data:
                try:
                    frame = pickle.loads(data)
                    cv2.imshow("Video", frame)
                    
                    key = cv2.waitKey(1) & 0xFF
                    if key == ord("q"):
                        break
                except:
                    pass
            
    # send voice data to server
    def send_voice_data_to_server(self):
        while True:
            data = self.voice_recording_stream.read(1024)
            self.voice_ws.send_binary(data)

    # send video data to server
    def send_video_data_to_server(self):
        while True:
            ret, data = self.video_playing_stream.read(1024)
            frame = pickle.dumps(data)
            self.video_ws.send_binary(frame)
            
client = Client()

# close output window
cv2.destroyAllWindows()
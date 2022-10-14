interface AudioSource {
  float Volume();
}

class Microphone implements AudioSource {
  float Volume() {
    return 30.0;
  }
}

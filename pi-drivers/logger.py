
class Logger:
    def log(self, prompt) -> None:
        print(prompt)

    def logWithLevel(self, level, prompt) -> None:
        print(f'[{level}]: {prompt}')



import discord
import random
import os
from dotenv import load_dotenv

intents = discord.Intents.default()
intents.message_content = True
client = discord.Client(intents=intents)

responses = [
    "Better be seen at oxford than caught at john's",
    "rather go to oxford than st johns",
    "watch your mouth before you get sent to St Johns",
    "when do we demolish st johns again?",
    "they say st john's was built to make the rest of cambridge look better.",
    "the best part of visiting st johns is when you leave",
]


@client.event
async def on_message(message: discord.Message):
    if message.author == client.user:
        return

    if "john" not in message.content.lower():
        return

    response = random.choice(responses)
    await message.reply(response)


if __name__ == "__main__":
    load_dotenv()
    token = os.getenv("DISCORD_TOKEN")
    if token is None:
        exit(1)
    client.run(token)

const OpenAI =require('openai');
const config= require('./config.json');
const apiKey= config.OPENAI_API_KEY;
const openai = new OpenAI({ apiKey: apiKey });

async function main() {
  const completion = await openai.chat.completions.create({
    messages: [{ role: "system", content: "how to plan for a holiday?" }],
    model: "gpt-3.5-turbo",
  });

  console.log(completion.choices[0]);
}

main();
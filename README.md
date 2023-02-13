# go-example-chatgpt-copilot
Golang project experimenting with [ChatGPT](https://openai.com/blog/chatgpt/) and [GitHub Copilot](https://github.com/features/copilot) tooling.

NOTICE: This project uses GPT-3 generated content and has notices on each source file impacted.

# ChatGPT Requests/Responses
See [docs/ChatGPT Network Logs/README.md](/docs/ChatGPT%20Network%20Logs/README.md) to view the the ChatGPT requests/responses for helping creating this repo/src code.
- Since the API is [limited to https://chat.openai.com at the moment of this writing](https://twitter.com/OpenAI/status/1615160228366147585?ref_src=twsrc%5Egoogle%7Ctwcamp%5Eserp%7Ctwgr%5Etweet).

# ChatGPT "Animal API" Project Documentation (Code Gen via ChatGPT)
See [src/README](/src/README.md) for the documentation of the example API that ChatGPT generated. 

## Thoughts/Notes
- Great at boilerplate/small questions
- When even minor complexity comes up, **it stumbles quickly**
- For instance I asked for polymorphism for dog and cat, and it generated the code find, but then implemented dog and cat completely incorrect
    - It "used" the struct, but none of the methods
- Doesn't really have "context" of previous code, often I would ask for additions or removal to a code piece, and instead ChatGPT would provide an almost completely different solution/logic
- Co-pilot -> Also takes items very literally, great for ease of use and unit testing, and documentation
  - Fails with complexity, but can at least encourage good naming/architecture (ex: new function/struct name)

### Swagger UI Rendering
https://ribice.medium.com/serve-swaggerui-within-your-golang-application-5486748a5ed4

Now instead of editing index.html, the file needing modification is [swagger-initializer.js](https://github.com/swagger-api/swagger-ui/blob/master/dist/swagger-initializer.js#L6)

## Summary
Early Stages, Great a Boilerplate. Great piece of tooling though, often it helped me "get" to the interesting logic.
- Needs: better context! Perhaps my requests weren't "worded" the best.
- It definitely **needs work w/ context/historical knowledge**, I struggled to get it to push out consistent code

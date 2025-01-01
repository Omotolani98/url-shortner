import openai
import re

# Replace with your OpenAI API key
openai.api_key = "your_key_here"


def generate_seo_text(topic):
    # Generate SEO content using OpenAI's Chat API
    messages = [
        {"role": "system", "content": "You are an SEO expert."},
        {"role": "user", "content": f"Generate a short SEO-friendly description and keywords for a webpage about {topic}."}
    ]

    response = openai.chat.completions.create(
        model="gpt-4",
        messages=messages,
        max_tokens=100,
        temperature=0.7
    )

    seo_text = response.choices[0].message.content.strip()
    return seo_text

def append_to_html(file_path, seo_text):
    # Read the existing HTML content
    with open(file_path, "r") as file:
        html_content = file.read()

    # Format the SEO content as HTML
    seo_content = f"""
    <div id='seo-content'>
        <strong>Description:</strong>
        {seo_text.split('Keywords:')[0].strip()}
        <br><br>
        <strong>Keywords:</strong>
        {seo_text.split('Keywords:')[1].strip()}
    </div>
    """

    # Insert SEO content below the form by targeting the end of the result div
    new_html_content = re.sub(
        r"(</div>\s*</div>\s*</body>)",
        f"{seo_content}\n</div></div></body>",
        html_content,
        count=1
    )

    # Write the modified HTML content back to the file
    with open(file_path, "w") as file:
        file.write(new_html_content)

def main():
    topic = "URL Shortener"
    seo_text = generate_seo_text(topic)
    print("Generated SEO Text:\n", seo_text)

    # Append SEO content to your HTML file
    html_file_path = "../internal/views/index.html"  # Replace with your actual file path
    append_to_html(html_file_path, seo_text)
    print(f"SEO content appended to {html_file_path}")

if __name__ == "__main__":
    main()
# **Quiz App**

This is a quiz app that allows the user to select a topic and answer a number of questions within that topic. The app will display the score at the end.

## **Dependencies**

- github.com/rama-kairi/quiz-app/data
- github.com/rama-kairi/quiz-app/helpers

## **Usage**

1. Run the **`main`** function.
2. Select a topic from the list of options.
3. Select the number of questions you want to answer.
4. Answer the questions by selecting the correct answer from the list of options.
5. The app will display your score at the end.

## **Customization**

You can customize the quiz data by modifying the **`quizData`** variable in the **`data`** package. Each element in the slice should be a map with the following keys:

- **`question`**: a string representing the question
- **`correct_answer`**: a string representing the correct answer

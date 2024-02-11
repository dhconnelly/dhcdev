=== Notes: Learning Machine Learning ===

[[home]](/)

# Notes: Learning Machine Learning

_February 1, 2023_

I've decided to spend some time digging into Machine Learning and Deep Learning.
I took an ML course in grad school but forgot most of it, and I worked on a
system for Google Maps back in 2015-2016 that predicated hundreds of categorical
labels for tens of millions of POIs, but I only understood the model at a
superficial level, and it was the days before TensorFlow, so we had to build a
lot of custom stuff. A lot has happened since then!

I'll link notes and code below as I go along.

***Edit***: This project is complete. See the [retrospective](#retrospective) below.

## Table of Contents

-   [Google's Machine Learning Crash Course](#google-s-machine-learning-crash-course)
-   [Hands-On Machine Learning with Scikit-Learn, Keras & TensorFlow](#hands-on-machine-learning-with-scikit-learn-keras-and-tensorflow) (abandoned)
-   [Independent practice](#independent-practice)
-   [Coursera Machine Learning Specialization](#coursera-machine-learning-specialization)
-   [Retrospective](#retrospective)

## Google's Machine Learning Crash Course

**Course**: [https://developers.google.com/machine-learning/crash-course](https://developers.google.com/machine-learning/crash-course)

**Status**: Done

**Review**:

This ~15 hour video-heavy course provides a broad overview of the field. The
emphasis is on linear and logistic regression, neural networks, and
applied/engineering factors. There are also some "programming exercises" that
provide ~zero learning value, as you do little more than just read and execute
existing TensorFlow code. But still this is a very good and information-dense
course that I would highly recommend as an introduction.

**Notes**:

1. [Framing](/mlcc/mlcc-1-framing.pdf)
2. [Descending into ML](/mlcc/mlcc-2-descending-into-ml.pdf)
3. [Reducing Loss](/mlcc/mlcc-3-reducing-loss.pdf)
4. [First Steps with TensorFlow](/mlcc/mlcc-4-first-steps-tf.pdf)
5. [Real Datasets](/mlcc/mlcc-5-real-dataset.pdf)
6. [Generalization](/mlcc/mlcc-6-generalization.pdf)
7. [Training, Validation, Test Sets](/mlcc/mlcc-7-train-validate-test.pdf)
8. [Representation](/mlcc/mlcc-8-representation.pdf)
9. [Feature Crosses](/mlcc/mlcc-9-feature-crosses.pdf)
10. [Regularization: Simplicity](/mlcc/mlcc-10-regularization-simplicity.pdf)
11. [Logistic Regression](/mlcc/mlcc-11-logistic-regression.pdf)
12. [Classification](/mlcc/mlcc-12-classification.pdf)
13. [Regularization: Sparsity](/mlcc/mlcc-13-regularization-sparsity.pdf)
14. [Neural Networks](/mlcc/mlcc-14-neural-nets.pdf)
15. [Training Neural Networks](/mlcc/mlcc-15-training-neural-nets.pdf)
16. [Multi-Class Neural Networks](/mlcc/mlcc-16-multi-class-nns.pdf)
17. [Embeddings](/mlcc/mlcc-17-embeddings.pdf)
18. [ML Engineering](/mlcc/mlcc-18-ml-eng.pdf)
19. [ML Fairness](/mlcc/mlcc-19-fairness.pdf)
20. [Real-World Examples](/mlcc/mlcc-20-examples.pdf)
21. [Guidelines](/mlcc/mlcc-21-guidelines.pdf)

## Hands-On Machine Learning with Scikit-Learn, Keras, and TensorFlow

**Book**: [https://www.oreilly.com/library/view/hands-on-machine-learning/9781098125967/](https://www.oreilly.com/library/view/hands-on-machine-learning/9781098125967/)

**Status**: Abandoned. Fantastic book but goes way deeper into implementation
concerns and library/tool usage than I was looking for at the time.

**Notes and Colabs**:

-   [Annotated Machine Learning Project Checklist](/ml-project-checklist.html) taken from Appendix A plus details and tips from other chapters
-   Chapter 2: End-to-End ML Project - [Colab](https://colab.research.google.com/drive/18JrgQZ3DRNn3qXmExQKAR9kuHNKOAdnn?usp=sharing)
-   Chapter 3: Classification - [Colab](https://colab.research.google.com/drive/1mhl0SME75Fsa9fd11hDHCYJ-8D0A9OjK)
-   Chapter 4: Training Models - [Colab](https://colab.research.google.com/drive/1YBZ9bZXJqzpxEEu9rNnRLZVaY3WeOdqC)

## Independent practice

-   Kaggle: [MNIST without using neural networks](https://www.kaggle.com/dhconnelly/mnist-without-neural-nets)
-   Kaggle: [Titanic](https://www.kaggle.com/code/dhconnelly/titanic)
-   Kaggle: [Spaceship Titanic](https://www.kaggle.com/code/dhconnelly/spaceship-titanic/notebook)
-   Kaggle: [House Prices](https://www.kaggle.com/code/dhconnelly/house-prices)
-   [Linear regression with California housing data](https://colab.research.google.com/drive/1zj7b3Bzh7T9HCPQDM90zL6J0IM22Dvoa?usp=sharing)
-   [Predicting hotel reservation cancellation with logistic regression](https://colab.research.google.com/drive/1-ixQMV5EwC7emLaUO9KN9oTMV1Oz8TC7#scrollTo=NFDcZ_FO01LX)
-   [MNIST with a Keras softmax neural network](https://colab.research.google.com/drive/1kPpnrJMVmmQ_tsnacmWSXSj9GdHPXG-H)
-   [Predicting wine quality with linear regression and cross-validation](https://colab.research.google.com/drive/1oZw2dA2rFpjYHeqgLm3Ebj2QS87_Al_s)
-   [Forest cover type classification with XGBoost](https://colab.research.google.com/drive/1tCq19-iw8tLTT_foiI1w_D0YMnaBPmAx?usp=sharing)
-   [Ad engagement regression with XGBoost](https://colab.research.google.com/drive/1mWdS4MMhtpvYqjK8vQcThXS7I90fqO8E)

## Coursera Machine Learning Specialization

**Course**: [https://www.deeplearning.ai/courses/machine-learning-specialization/](https://www.deeplearning.ai/courses/machine-learning-specialization/)

**Status**: Done

**Review**:

Great course for learning the algorithms and ideas in more detail than the
Google crash course. Lots of math (like backpropagation from scratch and so on)
helped me develop a better sense of how things work than just playing around
with Kaggle problems. Also, Andrew Ng's comments and tips based on his
extensive experience were a nice feature. Unfortunately the programming labs
were almost entirely useless - you implement a tiny bit of an algorithm based
on literally translating a couple of equations into Python. It would have been
much more interesting if the labs required you to build up the entire pipeline
and algorithm from scratch (or with hints). If I hadn't worked through the
first couple of chapters of
[HOML](#hands-on-machine-learning-with-scikit-learn-keras-and-tensorflow) I
wouldn't have had any idea how the labs worked.

**Notes and colabs**:

-   Course 1: Supervised Machine Learning: Regression and Classification
    -   Intro to Machine Learning: [notes](/coursera/coursera-ml-c1-wk1.pdf)
    -   Multiple Linear Regression: [notes](/coursera/coursera-ml-c1-wk2.pdf)
    -   Classification: [notes](/coursera/coursera-ml-c1-wk3.pdf)
-   Course 2: Advanced Learning Algorithms
    -   Neural Networks: [notes](/coursera/coursera-ml-c2-wk1.pdf)
    -   Training Neural Networks: [notes](/coursera/coursera-ml-c2-wk2.pdf)
    -   Applying Machine Learning: [notes](/coursera/coursera-ml-c2-wk3.pdf)
    -   Decision Trees: [notes](/coursera/coursera-ml-c2-wk4.pdf)
-   Course 3: Beyond supervised learning
    -   Unsupervised learning: [notes](/coursera/coursera-ml-c3-wk1.pdf)
    -   Recommender systems: [notes](/coursera/coursera-ml-c3-wk2.pdf)
    -   Reinforcement learning: [notes](/coursera/coursera-ml-c3-wk3.pdf)

## Retrospective

I spent a year on this project (Feb 2023 to Feb 2024), with almost all of the
work work happening in the first six months (before I relocated to France over
the summer, when I had new problems to deal with). The effort absolutely paid
off professionally. Even if I didn't transition to AI work, my [current
company](https://www.teads.com/) does a lot of ML, and the time I spent on this
material made it possible for me to understand what kind of approaches we use
and why and their tradeoffs, identify the largest ML engineering challenges
we're facing, and have reasonable conversations with the team members about
their work and its difficulty and impact. It also made it possible to reframe
my previous work on a production ML system at Google in 2015-2016 into modern
language and concepts.

My next step is to dig further into Deep Learning and building stuff. I'll
make a new post about that :)

[[home]](/)

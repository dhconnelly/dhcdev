=== Notes: Learning Machine Learning ===


[[home]](/)

# Notes: Learning Machine Learning

_February 1, 2023_

I've decided to spend some time digging into Machine Learning and Deep Learning.
I took an ML course in grad school but forgot most of it, and I worked on a
system for Google Maps that predicated hundreds of categorical labels for tens
of millions of POIs, but without really understanding the details and math
behind the model.

I'll link notes and notebooks below as I go along.

## Table of Contents

-   [Google's Machine Learning Crash Course](#google-s-machine-learning-crash-course)
-   [Coursera Machine Learning Specialization](#coursera-machine-learning-specialization)
-   [Interlude: Some Kaggle submissions](#kaggle-submissions)
-   [Hands-On Machine Learning with Scikit-Learn, Keras & TensorFlow](#hands-on-machine-learning-with-scikit-learn-keras-and-tensorflow)

## Google's Machine Learning Crash Course

**Course**: [https://developers.google.com/machine-learning/crash-course](https://developers.google.com/machine-learning/crash-course)

**Status**: Done

**Review**:

This ~15 hour video-heavy course provides a broad overview of the field. The
emphasis is on linear and logistic regression, neural networks, and
applied/engineering factors. There are also some "programming exercises" that
provide ~zero learning value, as you do little more than just read and execute
existing TensorFlow code.

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

## Coursera Machine Learning Specialization

**Course**: [https://www.deeplearning.ai/courses/machine-learning-specialization/](https://www.deeplearning.ai/courses/machine-learning-specialization/)

**Status**: Finished courses 1 and 2. Got bored and skipped course 3 (reinforcement learning). Moved on to HOML (see next section).

**Notes and colabs**:

-   Course 1: Supervised Machine Learning: Regression and Classification
    -   Intro to Machine Learning: [notes](/coursera/coursera-ml-c1-wk1.pdf)
    -   Multiple Linear Regression: [notes](/coursera/coursera-ml-c1-wk2.pdf)
    -   Independent practice: [Linear regression with California housing data](https://colab.research.google.com/drive/1zj7b3Bzh7T9HCPQDM90zL6J0IM22Dvoa?usp=sharing)
    -   Classification: [notes](/coursera/coursera-ml-c1-wk3.pdf)
    -   Independent practice: [Predicting hotel reservation cancellation with logistic regression](https://colab.research.google.com/drive/1-ixQMV5EwC7emLaUO9KN9oTMV1Oz8TC7#scrollTo=NFDcZ_FO01LX)
-   Course 2: Advanced Learning Algorithms
    -   Neural Networks: [notes](/coursera/coursera-ml-c2-wk1.pdf)
    -   Training Neural Networks: [notes](/coursera/coursera-ml-c2-wk2.pdf)
    -   Independent practice: [Handwritten digit classification with a Keras softmax neural network](https://colab.research.google.com/drive/1kPpnrJMVmmQ_tsnacmWSXSj9GdHPXG-H)
    -   Applying Machine Learning: [notes](/coursera/coursera-ml-c2-wk3.pdf)
    -   Independent practice: [Predicting wine quality with linear regression and cross-validation](https://colab.research.google.com/drive/1oZw2dA2rFpjYHeqgLm3Ebj2QS87_Al_s) (lol it performs terribly)
    -   Decision Trees: [notes](/coursera/coursera-ml-c2-wk4.pdf)
    -   Independent practice: [Forest cover type classification with XGBoost](https://colab.research.google.com/drive/1tCq19-iw8tLTT_foiI1w_D0YMnaBPmAx?usp=sharing)
    -   Independent practice: [Ad engagement regression with XGBoost](https://colab.research.google.com/drive/1mWdS4MMhtpvYqjK8vQcThXS7I90fqO8E)

## Kaggle submissions

-   MNIST without using neural networks: [accuracy 0.97082](https://www.kaggle.com/dhconnelly/mnist-without-neural-nets)
-   Titanic: [accuracy 0.77751](https://www.kaggle.com/code/dhconnelly/titanic)
-   Spaceship Titanic: [accuracy 0.79752](https://www.kaggle.com/code/dhconnelly/spaceship-titanic/notebook)
-   House Prices: [error 0.15108](https://www.kaggle.com/code/dhconnelly/house-prices)

## Hands-On Machine Learning with Scikit-Learn, Keras, and TensorFlow

**Book**: [https://www.oreilly.com/library/view/hands-on-machine-learning/9781098125967/](https://www.oreilly.com/library/view/hands-on-machine-learning/9781098125967/)

**Status**: Just started. Going to experiment with [JAX](https://jax.readthedocs.io/en/latest/notebooks/quickstart.html) and [Flax](https://flax.readthedocs.io/en/latest/) too for some sections and exercises.

**Notes and Colabs**:

TODO
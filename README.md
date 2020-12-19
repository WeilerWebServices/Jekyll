# **[Jekyll](https://jekyllrb.com)**

![Jekyll](Jekyll.png)

#### ***[Jekyll](https://jekyllrb.com)*** is a blog-aware, static site generator in Ruby.

---

## [Setup](https://jekyllrb.com/docs/step-by-step/01-setup/)

#### Welcome to Jekyll's step-by-step tutorial. The goal of this tutorial is to take you from having some front end web development experience to building your first Jekyll site from scratch --- not relying on the default gem-based theme. Let's get into it!

---

## [Installation](https://jekyllrb.com/docs/step-by-step/01-setup/#installation "Permalink")

***Jekyll*** is a Ruby program so you need to install Ruby on your machine to begin with. Head over to the [install guide](https://jekyllrb.com/docs/installation/) and follow the instructions for your operating system.

With Ruby setup you can install Jekyll by running the following in your terminal:

```
gem install jekyll bundler
```

To create a new `Gemfile` to list your project's dependencies run:

```
bundle init
```

Now edit the `Gemfile` and add jekyll as a dependency:

```
gem "jekyll"
```

Finally run `bundle` to install ***jekyll*** for your project.

You can now prefix all ***jekyll*** commands listed in this tutorial with `bundle exec` to make sure you use the jekyll version defined in your `Gemfile`.

---

## [Create a site](https://jekyllrb.com/docs/step-by-step/01-setup/#create-a-site "Permalink")

It's time to create a site! Create a new directory for your site, you can name it whatever you'd like. Through the rest of this I'll refer to this directory as "root".

If you're feeling adventurous, you can also initialize a Git repository here. One of the great things about ***Jekyll*** is there's no database. All content and site structure are files which a Git repository can version. Using a repository is completely optional but it's a great habit to get into. You can learn more about using Git by reading through the [Git Handbook](https://guides.github.com/introduction/git-handbook/).

Let's add your first file. Create `index.html` in the root with the following content:

```
<!doctype html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Home</title>
  </head>
  <body>
    <h1>Hello World!</h1>
  </body>
</html>
```

---

## [Build](https://jekyllrb.com/docs/step-by-step/01-setup/#build "Permalink")

***Jekyll*** is a static site generator so we need Jekyll to build the site before we can view it. There are two commands you can run in the root of your site to build it:

-   `jekyll build` - Builds the site and outputs a static site to a directory called `_site`.
-   `jekyll serve` - Does the same thing except it rebuilds any time you make a change and runs a local web server at `http://localhost:4000`.

When you're developing a site you'll use `jekyll serve` as it updates with any changes you make.

Run `jekyll serve` and go to [http://localhost:4000](http://localhost:4000/) in your browser. You should see "Hello World!".

---

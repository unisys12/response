---
title: "Introduction"
description: "Introduction to Response for administrators of a single Response instance or cluster."
lead: "Response is designed to be easy to deploy for all skill levels. This introduction will help familiarize you with the overall requirements for running Response."
date: 2020-10-06T08:48:57+00:00
lastmod: 2020-10-06T08:48:57+00:00
draft: false
images: []
menu:
  docs:
    parent: "admin"
weight: 100
toc: true
---

## Welcome

Welcome to Response and thanks for trying it out and/or adopting it at your community. We've put a lot of effort into making sure your
experience with Response as an administrator is as smooth as can be.

Response is still a young project and we expect our documentation to improve as the community begins to provide feedback on topics
covered. If you have feedback and/or suggestions, please share them.

## Scenarios

Response supports a few deployment scenarios, namely Single Instance and Cluster deployments.

### Single Instance

Single instance deployments are the simplest scenario and are recommended for smaller communities, when first deploying Response, or
when you've outgrown a single instance and need to balance the load across multiple instances.

As of right now, we have not performed real-world load testing on Response but do plan to in the future.

### Cluster

Cluster deployments are more complex and involve configuring more than one Response instance to communicate and find other
Response instances. In general, only large and workload-heavy deployments truly need to consider this scenario.

## Requirements

Response has very few external dependencies and for the most part bundles all of the necessary dependencies inside of the single
deployable binary. External dependencies do grow slightly when performing [Cluster deployments of Response](#cluster).

Response requires the following external dependencies in order to operate:

### Operating System

Response has full support for deployment on the following operating systems. The team recommends the use of Linux when possible
but understands that deployment on other operating systems may be necessary depending on knowledge of the administrators.

- Linux (Debian, Ubuntu, etc.)
- Windows
- macOS (both Intel and M1 chipsets are supported)

###

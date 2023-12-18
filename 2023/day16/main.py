def parse_input():
    with open("input.txt") as f:
        lines = f.read().splitlines()
    return lines

def energize(grid: list[str]):
    starting_item = grid[0][0]
    dir_map = {
        "curr_direction": "right",
        "curr_y": 0,
        "curr_x": 0,
        "curr_item": starting_item
    }
    queue = []
    visited = []
    queue.append(dir_map)

    while True:
        # Grab the item at the front of the queue
        curr_dir = queue.pop(0)
        
        # Pull out the keys from the curr_dir
        curr_direction = curr_dir["curr_direction"]
        curr_x = curr_dir["curr_x"]
        curr_y = curr_dir["curr_y"]
        curr_item = curr_dir["curr_item"]

        # If the we are heading right and we aren't at the end of the row
        if curr_direction == "right" and curr_x < len(grid[0])-1:
            curr_x += 1
            curr_item = grid[curr_y][curr_x]
            
            # If we have already visited this position, we don't want to add it to the queue again
            obj = { "curr_direction": curr_direction, "curr_x": curr_x, "curr_y": curr_y }
            if obj not in visited:
                visited.append(obj)
            else:
                continue
            # Based on the curr_item, we need to decide which direction to go next
            if curr_item == "|":
                # If we run into the pipe, and we are heading right, we need to add both "up" and "down" to the queue
                # along with the new position
                queue.append({
                    "curr_direction": "up",
                    "curr_x": curr_x,
                    "curr_y": curr_y,
                    "curr_item": curr_item
                })
                queue.append({
                    "curr_direction": "down",
                    "curr_x": curr_x,
                    "curr_y": curr_y,
                    "curr_item": curr_item
                })
            elif curr_item == "-":
                # If we run into a dash and we are heading right, we can continue through it in the same direction.
                # So we want to add the new position to the queue
                queue.append({
                    "curr_direction": "right",
                    "curr_x": curr_x,
                    "curr_y": curr_y,
                    "curr_item": curr_item
                })
            elif curr_item == '/':
                # If we run into a slash and we are heading right, we need to change directions to "up"
                queue.append({
                    "curr_direction": "up",
                    "curr_x": curr_x,
                    "curr_y": curr_y,
                    "curr_item": curr_item
                })
            elif curr_item == '\\':
                # If we run into a backslash and we are heading right, we need to change directions to "down"
                queue.append({
                    "curr_direction": "down",
                    "curr_x": curr_x,
                    "curr_y": curr_y,
                    "curr_item": curr_item
                })
            elif curr_item == '.':
                # We continue on in the direction we are heading.
                queue.append({
                    "curr_direction": "right",
                    "curr_x": curr_x,
                    "curr_y": curr_y,
                    "curr_item": curr_item
                })
        # If we are heading left and we aren't at the beginning of the row
        elif curr_direction == "left" and curr_x > 0:
            curr_x -= 1
            curr_item = grid[curr_y][curr_x]
            # If we have already visited this position, we don't want to add it to the queue again
            obj = { "curr_direction": curr_direction, "curr_x": curr_x, "curr_y": curr_y }
            if obj not in visited:
                visited.append(obj)
            else:
                continue
            # Based on the curr_item, we need to decide which direction to go next
            if curr_item == "|":
                # If we run into the pipe, and we are heading left, we need to add both "up" and "down" to the queue
                # along with the new position
                queue.append({
                    "curr_direction": "up",
                    "curr_x": curr_x,
                    "curr_y": curr_y,
                    "curr_item": curr_item
                })
                queue.append({
                    "curr_direction": "down",
                    "curr_x": curr_x,
                    "curr_y": curr_y,
                    "curr_item": curr_item
                })
            elif curr_item == "-":
                # If we run into a dash and we are heading left, we can continue through it in the same direction.
                # So we want to add the new position to the queue
                queue.append({
                    "curr_direction": "left",
                    "curr_x": curr_x,
                    "curr_y": curr_y,
                    "curr_item": curr_item
                })
            elif curr_item == '/':
                # If we run into a slash and we are heading left, we need to change directions to "down"
                queue.append({
                    "curr_direction": "down",
                    "curr_x": curr_x,
                    "curr_y": curr_y,
                    "curr_item": curr_item
                })
            elif curr_item == '\\':
                # If we run into a backslash and we are heading left, we need to change directions to "up"
                queue.append({
                    "curr_direction": "up",
                    "curr_x": curr_x,
                    "curr_y": curr_y,
                    "curr_item": curr_item
                })
            elif curr_item == '.':
                # We continue on in the direction we are heading.
                queue.append({
                    "curr_direction": "left",
                    "curr_x": curr_x,
                    "curr_y": curr_y,
                    "curr_item": curr_item
                })
        # If we are heading up and we aren't at the beginning of the column
        elif curr_direction == "up" and curr_y > 0:
            curr_y -= 1
            curr_item = grid[curr_y][curr_x]
            # If we have already visited this position, we don't want to add it to the queue again
            obj = { "curr_direction": curr_direction, "curr_x": curr_x, "curr_y": curr_y }
            if obj not in visited:
                visited.append(obj)
            else:
                continue
            # Based on the curr_item, we need to decide which direction to go next
            if curr_item == "|":
                # If we run into the pipe, and we are heading up, we can continue through it in the same direction.
                # So we want to add the new position to the queue
                queue.append({
                    "curr_direction": "up",
                    "curr_x": curr_x,
                    "curr_y": curr_y,
                    "curr_item": curr_item
                })
            elif curr_item == "-":
                # If we run into a dash and we are heading up, we need to change directions to both "left" and "right"
                queue.append({
                    "curr_direction": "right",
                    "curr_x": curr_x,
                    "curr_y": curr_y,
                    "curr_item": curr_item
                })
                queue.append({
                    "curr_direction": "left",
                    "curr_x": curr_x,
                    "curr_y": curr_y,
                    "curr_item": curr_item
                })
            elif curr_item == '/':
                # If we run into a slash and we are heading up, we need to change directions to "right"
                queue.append({
                    "curr_direction": "right",
                    "curr_x": curr_x,
                    "curr_y": curr_y,
                    "curr_item": curr_item
                })
            elif curr_item == '\\':
                # If we run into a backslash and we are heading up, we need to change directions to "left"
                queue.append({
                    "curr_direction": "left",
                    "curr_x": curr_x,
                    "curr_y": curr_y,
                    "curr_item": curr_item
                })
            elif curr_item == '.':
                # We continue on in the direction we are heading.
                queue.append({
                    "curr_direction": "up",
                    "curr_x": curr_x,
                    "curr_y": curr_y,
                    "curr_item": curr_item
                })
        # If we are heading down and we aren't at the end of the column
        elif curr_direction == "down" and curr_y < len(grid)-1:
            curr_y += 1
            curr_item = grid[curr_y][curr_x]
            # If we have already visited this position, we don't want to add it to the queue again
            obj = { "curr_direction": curr_direction, "curr_x": curr_x, "curr_y": curr_y }
            if obj not in visited:
                visited.append(obj)
            else:
                continue
            # Based on the curr_item, we need to decide which direction to go next
            if curr_item == "|":
                # If we run into the pipe, and we are heading down, we can continue through it in the same direction.
                # So we want to add the new position to the queue
                queue.append({
                    "curr_direction": "down",
                    "curr_x": curr_x,
                    "curr_y": curr_y,
                    "curr_item": curr_item
                })
            elif curr_item == "-":
                # If we run into a dash and we are heading down, we need to change directions to both "left" and "right"
                queue.append({
                    "curr_direction": "right",
                    "curr_x": curr_x,
                    "curr_y": curr_y,
                    "curr_item": curr_item
                })
                queue.append({
                    "curr_direction": "left",
                    "curr_x": curr_x,
                    "curr_y": curr_y,
                    "curr_item": curr_item
                })
                # print(queue)
            elif curr_item == '/':
                # If we run into a slash and we are heading down, we need to change directions to "left"
                queue.append({
                    "curr_direction": "left",
                    "curr_x": curr_x,
                    "curr_y": curr_y,
                    "curr_item": curr_item
                })
            elif curr_item == '\\':
                # If we run into a backslash and we are heading down, we need to change directions to "right"
                queue.append({
                    "curr_direction": "right",
                    "curr_x": curr_x,
                    "curr_y": curr_y,
                    "curr_item": curr_item
                })
            elif curr_item == '.':
                # We continue on in the direction we are heading.
                queue.append({
                    "curr_direction": "down",
                    "curr_x": curr_x,
                    "curr_y": curr_y,
                    "curr_item": curr_item
                })

        # If we have visited the same position twice, we have found the loop
        if len(visited) > 1 and visited[0].get("curr_x") == visited[-1].get("curr_x") and visited[0].get("curr_y") == visited[-1].get("curr_y") and visited[0].get("curr_direction") == opposite(visited[-1].get("curr_direction")):
            break
        
    return visited

def opposite(dir: str):
    if dir == "up":
        return "down"
    elif dir == "down":
        return "up"
    elif dir == "left":
        return "right"
    elif dir == "right":
        return "left"


def part1():
    input = parse_input()
    grid = []
    for line in input:
        grid.append(list(line))
    path = energize(grid)
    print(path)
part1()
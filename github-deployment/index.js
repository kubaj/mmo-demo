/**
 * HTTP Cloud Function simulating GitHub webhook receiver.
 *
 * @param {Object} req Cloud Function request context.
 * @param {Object} res Cloud Function response context.
 */
exports.githubWH = (req, res) => {
    console.log(req.body);
    res.send('OK');
};